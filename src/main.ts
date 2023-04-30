import death from "death";
import { environment } from "./lib/environment";
import { parseJWTMiddleware } from "./middleware/parse-jwt";
import { bitmaskMiddleware } from "./middleware/bitmask-middleware";
import { API_PERMISSIONS } from "./permissions";
import express, { type Request, type Response } from "express";
import { PrismaClient } from "@prisma/client";
import bodyParser from "body-parser";
import { z } from "zod";
const prisma = new PrismaClient();
const app = express();
const { PORT } = environment;
const server = app.listen( `${PORT}`, () => {
    console.log( `server started at http://localhost:${ PORT }` );
});

app.use(bodyParser.json({
  strict: true
}));

app.use(parseJWTMiddleware);

app.get("/minecraft/:uuid/authorizations",bitmaskMiddleware(API_PERMISSIONS.SERVER_AUTHORIZATION_VIEW), async (request: Request, response: Response) => {
  const { uuid } = request.params;
  try {
    const minecraftAccount = await prisma.minecraftAccount.findFirst({
      where: {
        id: uuid
      },
      select: {
        owner: true
      }
    });
    if(!minecraftAccount) {
      response.status(404).json({
        error: "minecraft account not found"
      })
      return;
    }
    const authorizations = await prisma.serverAuthorization.findMany({
      where: {
        userId: minecraftAccount.owner.id
      }
    });
    response.json(authorizations);
  } catch(error) {
    console.error(error);
    response.status(500).json({
      error: "internal server error"
    });
  }
});

const authorizationPostSchema = z.object({
  serverId: z.string(),
  userId: z.string()
})

type AuthorizationPost =  typeof authorizationPostSchema._type;

app.post("user/:userId/authorizations",bitmaskMiddleware(API_PERMISSIONS.SERVER_AUTHORIZATION_GRANT), async (request: Request, response: Response) => {
  let authorization: AuthorizationPost
  try {
    authorization = await authorizationPostSchema.parseAsync(request.body)
  } catch(error) {
    response.status(400).json(error);
    return;
  }
  const { serverId, userId } = authorization;
  try {
    const exists = await prisma.serverAuthorization.findFirst({
      where: {
        serverId,
        userId
      }
    });
    if(exists) {
      response.status(409).json({
        error: "authorization already exists"
      })
      return;
    }
    const authorizedUser = await prisma.serverAuthorization.create({
      data: {
        serverId,
        userId
      }
    });
    if(authorizedUser) {
      response.json(authorizedUser)
      return;
    }
  } catch(error) {
    console.error(error);
    response.status(500).json({
      error: "internal server error"
    });
  }
});

const minecraftAccountPostSchema = z.object({
  uuid: z.string().min(36)
})

type MinecraftAccountPost = typeof minecraftAccountPostSchema._type;

app.post("/user/:userId/minecraft/account",bitmaskMiddleware(API_PERMISSIONS.SERVER_AUTHORIZATION_GRANT), async (request: Request, response: Response) => {
  let minecraftAccount: MinecraftAccountPost
  try {
    minecraftAccount = await minecraftAccountPostSchema.parseAsync(request.body)
  } catch(error) {
    response.status(400).json(error);
    return;
  }
  const { userId } = request.body
  const { uuid } = minecraftAccount;
  try {
    const exists = await prisma.user.findFirst({
      where: {
        id: userId
      }
    });
    if(!exists) {
      response.status(404).json({
        error: "user does not exist"
      })
      return;
    }
    const existsMinecraftAccount = await prisma.minecraftAccount.findFirst({
      where: {
        id: uuid
      }
    });
    if(existsMinecraftAccount) {
      response.status(409).json({
        error: "minecraft account already exists"
      })
      return;
    }
    const account = await prisma.minecraftAccount.create({
      data: {
        id: uuid,
        ownerId: exists.id,
      }
    })
    response.status(201).json(account);
  }
  catch(error) {
    console.error(error)
    response.status(500).json({
      error: "internal server error"
    })
  }
});


death(() => {
  server.close()
})
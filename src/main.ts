import death from "death";
import { environment } from "./lib/environment";
import { parseJWTMiddleware } from "./middleware/parse-jwt";
import { useRequestContext } from "./lib/context";
import { bitmaskMiddleware } from "./middleware/bitmask-middleware";
import { API_PERMISSIONS } from "./permissions";
import express, { type Request, type Response } from "express";
import { PrismaClient } from "@prisma/client";
const prisma = new PrismaClient();
const app = express();
const { PORT } = environment;
const server = app.listen( `${PORT}`, () => {
    console.log( `server started at http://localhost:${ PORT }` );
});
app.use(parseJWTMiddleware)
app.get('/', (request, response) => {
  response.send('Hello World!')
  const permissions = useRequestContext<number>(request, "permissions");
  console.log(permissions)
});

app.get("/servers/:serverId/authorizations/user/:userId",bitmaskMiddleware(API_PERMISSIONS.SERVER_AUTHORIZATION_VIEW), async (request: Request, response: Response) => {
  const { serverId, userId } = request.params;
  const authorizedUser = await prisma.serverAuthorization.findFirst({
    where: {
      serverId,
      userId
    }
  });
  if(authorizedUser) {
    response.json(authorizedUser)
    return;
  }
  response.status(404).json({
    error: "user not found"
  });
});

app.post("/servers/:serverId/authorizations/user/:userId",bitmaskMiddleware(API_PERMISSIONS.SERVER_AUTHORIZATION_GRANT), async (request: Request, response: Response) => {
  const { serverId, userId } = request.params;
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


death(() => {
  server.close()
})
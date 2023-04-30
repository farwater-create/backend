import jwt from "jsonwebtoken";
import { environment } from "../lib/environment";
import { type NextFunction, type Request, type Response } from "express";
import { setRequestContext } from "../lib/context";
import { jwtSchema } from "../lib/jwt";
const { JWT_SIGNING_KEY } = environment;
export const parseJWTMiddleware = (request: Request, response: Response, next: NextFunction): void => {
    const token = request.headers.authorization;
    if(token === undefined) {
        response.status(401).json({
            error: "unauthorized"
        });
        return;
    }
    try {
        const decoded = jwt.verify(token, JWT_SIGNING_KEY);
        const parsed = jwtSchema.parse(decoded)
        setRequestContext(request, "permissions", parsed.permissions);
    } catch(error) {
        console.error(error);
        response.status(401).json({
            error: "unauthorized"
        });
        return;
    }
    next();
}
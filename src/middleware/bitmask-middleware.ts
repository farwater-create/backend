import { type NextFunction, type Request, type Response } from "express";
import { useRequestContext } from "../lib/context";

export const bitmaskMiddleware = (mask: number) => {
    return (request: Request, response: Response, next: NextFunction) => {
        const permissions = useRequestContext<number>(request, "permissions");
        const result = mask & permissions;
        if(result === mask) {
            next()
            return;
        }
        response.status(401).json({
            error: `missing permissions: ${mask} have: ${permissions}`
        })
    }
}
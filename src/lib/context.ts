import { type Request } from "express";

export const requestContext = new WeakMap<Request, Map<string, unknown>>();

export const useRequestContext = <I>(request: Request, key: string): I => {
    return requestContext.get(request)?.get(key) as I;
}

export const setRequestContext = <I>(request: Request, key: string, value: I): I => {
    if(!requestContext.has(request)) requestContext.set(request, new Map())
    return requestContext.get(request)?.set(key, value) as I;
}
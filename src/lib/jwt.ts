import z from "zod";

export const jwtSchema = z.object({
    permissions: z.number().min(0)
})

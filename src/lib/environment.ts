import z from "zod";
import { config } from "dotenv";
config()
const environmentSchema = z.object({
    JWT_SIGNING_KEY: z.string().min(24),
    PORT: z.string(),
    DATABASE_URL: z.string()
});

export const environment = environmentSchema.parse(process.env);

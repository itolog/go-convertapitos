import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

export const formSchema = toTypedSchema(
  z.object({
    email: z.string().email().min(1),
    name: z.string().min(1).max(70),
    verifiedEmail: z.boolean(),
    roleId: z.string().min(1),
  }),
);

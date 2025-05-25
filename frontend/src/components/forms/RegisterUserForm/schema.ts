import { toTypedSchema } from "@vee-validate/zod";
import * as z from "zod";

const formSchema = toTypedSchema(
  z.object({
    name: z.string().min(1).max(70),
    email: z.string().email().min(1),
    password: z.string().min(6).max(128),
  }),
);

export default formSchema;

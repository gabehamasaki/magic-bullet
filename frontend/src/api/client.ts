import createClient from "openapi-fetch";
import type { paths } from "./schema";

const baseUrl = import.meta.env.API_BASE_URL || "http://localhost:8080/api/v1";

export const apiClient = createClient<paths>({
  baseUrl,
});

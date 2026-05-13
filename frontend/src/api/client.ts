import createClient from "openapi-fetch";
import type { paths } from "./schema";

const runtimeApiBaseUrl = typeof process !== "undefined" ? process.env.API_BASE_URL : undefined;

const baseUrl = runtimeApiBaseUrl || import.meta.env.PUBLIC_API_BASE_URL || "http://localhost:8080/api/v1";

export const apiClient = createClient<paths>({
  baseUrl,
});

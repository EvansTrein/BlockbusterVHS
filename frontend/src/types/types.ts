export type HttpMethod = "GET" | "POST" | "PUT" | "DELETE";

export interface Parameter {
  title: string;
  count: number;
}

export interface Operation {
  title: string;
  description: string;
  apiPath: string;
  method: HttpMethod;
}

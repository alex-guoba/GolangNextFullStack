import { env } from "@/env"
import { ReqArticleList, ReqTagList, RspArticleList, RspTagList } from "./types";

type QueryParamValue = string | number | boolean | null | undefined;
type QueryParams<T> = { [K in keyof T]: T[K] extends QueryParamValue ? T[K] : never };

const getQueryURL = <T extends QueryParams<T>>(path: string, query?: T): string => {
  const baseUrl = `${env.BACKEND_URL}/api/v1/${path}`;
  
  const url = new URL(baseUrl);
 
  if (query) {
    const params = new URLSearchParams();
    for (const [key, value] of Object.entries(query)) {
      if (value !== null && value !== undefined) {
        params.append(key, value.toString());
      }
    }
    url.search = params.toString();
  }
 
  // 返回构建好的 URL
  return url.toString();
};
 

export const fetchArticles = async (req: ReqArticleList): Promise<RspArticleList> => {
  const articles = await fetch(getQueryURL<ReqArticleList>("articles", req),  {
    next: {
      revalidate: 0,
    }
  });
  return articles.json() as Promise<RspArticleList>;
}


export const fetchTags = async (req: ReqTagList): Promise<RspTagList> => {
  const articles = await fetch(getQueryURL<ReqTagList>("tags", req),  {
    next: {
      revalidate: 0,
    }
  });
  return articles.json() as Promise<RspTagList>;
}
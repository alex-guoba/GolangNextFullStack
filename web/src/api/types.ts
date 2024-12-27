
export interface Tag {
  id: number;
  created_by: string;
  modified_by: string;
  is_del: number;
  name: string;
  state: number;
}
 
export interface Article {
  id: number;
  title: string;
  desc: string;
  content: string;
  cover_image_url: string;
  state: number;
  tags: Tag[];
}
 
export interface Pager {
  page: number;
  page_size: number;
  total_rows: number;
}
 
export interface RspArticleList {
  code: number;
  msg: string;
  data: {
      list?: Article[];
      pager: Pager;
  };
};

export interface ReqPager {
  page?: number;
  page_size?: number;
};

export interface ReqArticleList extends ReqPager {
  tag_id?: number;
  state?: number;
};

export interface RspTagList {
  code: number;
  msg: string;
  data: {
      list?: Tag[];
      pager: Pager;
  };
};

export interface ReqTagList extends ReqPager {
  name?: string;
  state?: number;
};


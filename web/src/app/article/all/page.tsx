import Shell from "@/components/shell/shell";
import {
  PageHeader,
  PageHeaderHeading,
  PageHeaderDescription,
} from "@/components/page-header";
import { fetchArticles } from "@/api/articles";
import { RspArticleList } from "@/api/types";
import ArticleList from "@/components/article/article-list";
import { ArticlePagination } from "@/components/article/pagination";

type SearchParams = Promise<{ [key: string]: string | string[] | undefined }>;

async function ArticleRender({
  lst,
  tag_id,
}: {
  lst: RspArticleList;
  tag_id?: number;
}) {
  if (!lst || lst.code != 0) {
    throw new Error("Query Error");
  }

  if (lst.data && lst.data.list) {
    return (
      <>
        <ArticleList articles={lst.data.list} tag_id={tag_id} />
        <ArticlePagination
          total={lst.data.pager.total_rows}
          pageSize={lst.data.pager.page_size}
        ></ArticlePagination>
      </>
    );
  }

  return <div>No Articles</div>;
}

export default async function All(props: { searchParams: SearchParams }) {
  // https://nextjs.org/docs/app/building-your-application/upgrading/version-15#params--searchparams
  const sp = await props.searchParams;

  const page = Number(sp["page"]) || 1;
  const tag_id = Number(sp["tag_id"]) || 0;
  const articles = await fetchArticles({
    page,
    tag_id,
  });

  return (
    <Shell variant="default" className="md:pb-10">
      <PageHeader>
        <PageHeaderHeading size="lg" className="">
          All Articles
        </PageHeaderHeading>
        <PageHeaderDescription size="xs" className="py-4">
          all articles of this blog
        </PageHeaderDescription>
        <ArticleRender lst={articles} tag_id={tag_id} />
      </PageHeader>
    </Shell>
  );
}

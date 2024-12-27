import Shell from "@/components/shell/shell";
import {
  PageHeader,
  PageHeaderHeading,
  PageHeaderDescription,
} from "@/components/page-header";
import { fetchTags } from "@/api/articles";
import { RspTagList } from "@/api/types";
import { ArticlePagination } from "@/components/article/pagination";
import TagList from "@/components/article/tag-list";

type SearchParams = Promise<{ [key: string]: string | string[] | undefined }>;

async function TagsRender({ lst, name }: { lst: RspTagList; name?: string }) {
  if (!lst || lst.code != 0) {
    throw new Error("Query Error");
  }

  if (lst.data && lst.data.list) {
    return (
      <>
        <TagList tags={lst.data.list} name={name} />
        <ArticlePagination
          total={lst.data.pager.total_rows}
          pageSize={lst.data.pager.page_size}
        ></ArticlePagination>
      </>
    );
  }

  return <div>No tags</div>;
}

export default async function Tags(props: { searchParams: SearchParams }) {
  const sp = await props.searchParams;

  const page = Number(sp["page"]) || 1;
  const name = String(sp["name"] || "");
  const tags = await fetchTags({
    page,
    name,
  });
  return (
    <Shell variant="default" className="md:pb-10">
      <PageHeader>
        <PageHeaderHeading size="lg" className="text-center">
          Tags
        </PageHeaderHeading>
        <PageHeaderDescription size="xs" className="py-4">
          All the tags.
        </PageHeaderDescription>
        <TagsRender lst={tags} name={name} />
      </PageHeader>
    </Shell>
  );
}

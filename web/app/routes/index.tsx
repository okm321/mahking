import { createFileRoute } from "@tanstack/react-router";
import { HomeView } from "~/views/HomeView";

export const Route = createFileRoute("/")({
  head: () => ({
    meta: [
      { title: "Mahking | 麻雀の対局記録アプリ" },
      { name: "description", content: "Welcome to React Router!" },
    ],
  }),
  component: Home,
});

function Home() {
  return <HomeView />;
}

import * as React from "react";
import {
  HeadContent,
  Outlet,
  Scripts,
  createRootRoute,
} from "@tanstack/react-router";
import type { ErrorComponentProps } from "@tanstack/react-router";
import "../app.scss";
import { Header } from "~/components/Header";
import { Main } from "~/components/Main";

export const Route = createRootRoute({
  head: () => ({
    meta: [
      { charSet: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { title: "Mahking | 麻雀の対局記録アプリ" },
    ],
    links: [
      { rel: "preconnect", href: "https://fonts.googleapis.com" },
      {
        rel: "preconnect",
        href: "https://fonts.gstatic.com",
        crossOrigin: "anonymous",
      },
      {
        rel: "stylesheet",
        href: "https://fonts.googleapis.com/css2?family=Paytone+One&family=M+PLUS+Rounded+1c:wght@500;700&display=swap",
      },
      {
        rel: "preload",
        href: "/pinzu1-white.svg",
        as: "image",
      },
    ],
  }),
  shellComponent: RootDocument,
  component: RootComponent,
  errorComponent: DefaultErrorComponent,
});

function RootComponent() {
  return (
    <>
      <Header />
      <Main>
        <Outlet />
      </Main>
    </>
  );
}

function DefaultErrorComponent({ error }: ErrorComponentProps) {
  const details =
    import.meta.env.DEV && error instanceof Error
      ? error.message
      : "An unexpected error occurred.";
  const stack =
    import.meta.env.DEV && error instanceof Error ? error.stack : undefined;

  return (
    <main className="pt-16 p-4 container mx-auto">
      <h1>Error</h1>
      <p>{details}</p>
      {stack && (
        <pre className="w-full p-4 overflow-x-auto">
          <code>{stack}</code>
        </pre>
      )}
    </main>
  );
}

function RootDocument({ children }: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="ja">
      <head>
        <HeadContent />
      </head>
      <body>
        {children}
        <Scripts />
      </body>
    </html>
  );
}

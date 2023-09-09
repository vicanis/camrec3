import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import "./index.css";
import PageIndex from "./pages/index";
import PageDashboard from "./pages/dashboard";
import DeferredPageEvents, { LoadEvents } from "./pages/events";
import PageLive from "./pages/live";

const router = createBrowserRouter([
    {
        path: "/",
        element: <PageIndex />,
        children: [
            {
                index: true,
                element: <PageDashboard />,
            },
            {
                path: "events",
                element: <DeferredPageEvents />,
                loader: LoadEvents,
            },
            {
                path: "live",
                element: <PageLive />,
            },
        ],
    },
]);

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
    <React.StrictMode>
        <RouterProvider router={router} />
    </React.StrictMode>
);

import React, { ReactNode } from "react";
import { Await, useAsyncError, useLoaderData } from "react-router-dom";
import Loader from "./loader";
import Blur from "./blur";

export default function LoadablePage({
    renderer,
}: {
    renderer: (data: any) => ReactNode;
}) {
    const { data } = useLoaderData() as { data: any };

    return (
        <React.Suspense
            fallback={
                <Blur>
                    <Loader>Загрузка данных ...</Loader>
                </Blur>
            }
        >
            <Await resolve={data} errorElement={<ErrorElement />}>
                {renderer}
            </Await>
        </React.Suspense>
    );
}

function ErrorElement() {
    const error = useAsyncError() as Error;
    return <div>Ошибка загрузки данных: {error.message}</div>;
}

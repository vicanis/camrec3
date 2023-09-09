import { LoaderFunctionArgs, defer, useParams } from "react-router-dom";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import "dayjs/locale/ru";
import LoadablePage from "../components/loadablepage";
import { Event, EventList } from "../types/event";
import Icon from "@mdi/react";
import { mdiArrowLeft, mdiArrowRight } from "@mdi/js";

dayjs.extend(relativeTime);
dayjs.locale("ru");

export default function DeferredPageEvents() {
    return (
        <LoadablePage
            renderer={(data: EventList) => <PageEvents {...data} />}
        />
    );
}

function PageEvents({ count, items }: EventList) {
    const { day = dayjs().format("YYYYMMDD") } = useParams();

    const selectedDay = dayjs(day, "YYYYMMDD");

    return (
        <div className="my-3">
            <div className="flex gap-3">
                <div>
                    Всего событий: <b>{count}</b>
                </div>

                <div className="flex gap-3">
                    <a
                        href={
                            "/events/" +
                            dayjs(day).add(-1, "day").format("YYYYMMDD")
                        }
                    >
                        <Icon path={mdiArrowLeft} size={1} />
                    </a>
                    <div>{selectedDay.format("D MMMM YYYY")}</div>
                    <a
                        href={
                            "/events/" +
                            dayjs(day).add(1, "day").format("YYYYMMDD")
                        }
                    >
                        <Icon path={mdiArrowRight} size={1} />
                    </a>
                </div>
            </div>

            <div className="grid">
                {items.map((item) => (
                    <PageEvent key={item.id} {...item} />
                ))}
            </div>
        </div>
    );
}

function PageEvent({ id, timestamp, processed }: Event) {
    return (
        <div key={id} className="grid grid-cols-3 my-2">
            <div>{id}</div>
            <div>
                {dayjs(timestamp).format("DD.MM.YYYY HH:mm:ss Z")} (
                {dayjs(timestamp).fromNow()}){dayjs(timestamp).unix()}
            </div>
            <div>{processed ? "Обработано" : "Не обработано"}</div>
        </div>
    );
}

export function LoadEvents(args: LoaderFunctionArgs) {
    return defer({
        data: loader(args),
    });
}

async function loader({ params }: LoaderFunctionArgs) {
    const { day = dayjs().format("YYYYMMDD") } = params;

    const resp = await fetch(`/api/events?day=${day}`);
    const data = await resp.json();

    return data;
}

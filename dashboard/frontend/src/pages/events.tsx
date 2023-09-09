import { defer } from "react-router-dom";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import "dayjs/locale/ru";
import LoadablePage from "../components/loadablepage";
import { Event, EventList } from "../types/event";

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
    return (
        <div className="my-3">
            <div>Событий: {count}</div>

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
        <div key={id} className="grid grid-cols-2 my-2">
            <div>
                {dayjs(timestamp).format("DD.MM.YYYY HH:mm:ss Z")} (
                {dayjs(timestamp).fromNow()})
            </div>
            <div>{processed ? "Обработано" : "Не обработано"}</div>
        </div>
    );
}

export function LoadEvents() {
    return defer({
        data: loader(),
    });
}

async function loader() {
    const resp = await fetch("/api/events");
    const data = await resp.json();

    return data;
}

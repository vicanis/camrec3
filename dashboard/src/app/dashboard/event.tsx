import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
dayjs.extend(relativeTime);
import Video from "./video";

export default function EventList({ list }: { list: EventData[] }) {
    return (
        <ul className="grid gap-4">
            {list.map((item, index) => (
                <li key={item.id}>
                    <Event {...item} />
                </li>
            ))}
        </ul>
    );
}

function Event(e: EventData) {
    return (
        <div className="relative">
            <div className="absolute left-2 top-2 text-white">
                <div>{dayjs(e.timestamp).format("HH:mm:ss")}</div>
                <div>{dayjs(e.timestamp).fromNow()}</div>
            </div>
            <Video
                ts={dayjs(e.timestamp).format("YYYYMMDDHHmmss")}
                file={e.file}
            />
        </div>
    );
}

export type EventData = {
    id: string;
    timestamp: string;
    processed: boolean;
    file: string;
};

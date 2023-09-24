import dayjs, { Dayjs } from "dayjs";
import apicall from "@/apicall";
import EventList, { EventData } from "./event";
import Link from "next/link";

export default async function Page({
    searchParams,
}: {
    searchParams: { date: string | undefined };
}) {
    const date = dayjs(searchParams.date, "YYYYMM");
    const { events } = await getData(date);

    return (
        <div>
            <div
                className="fixed top-0 max-w-3xl mx-auto z-10 flex items-center justify-between w-full p-2"
                style={{
                    backgroundColor: "rgb(var(--background))",
                    color: "rgb(var(--foreground))",
                }}
            >
                <div>{events.length > 0 ? events.length : "No"} events</div>

                <div className="grid gap-2">
                    <div className="flex items-center gap-6 mx-auto w-max">
                        <Link
                            href={
                                "/dashboard/?date=" +
                                date.subtract(1, "day").format("YYYYMMDD")
                            }
                        >
                            <button
                                className="border-2 rounded-sm px-2"
                                style={{
                                    background: "rgb(var(--background))",
                                }}
                            >
                                {"<-"}
                            </button>
                        </Link>

                        <div>{dayjs(date).format("DD MMMM YYYY")}</div>

                        {!date.isSame(dayjs(), "date") && (
                            <Link
                                href={
                                    "/dashboard/?date=" +
                                    date.add(1, "day").format("YYYYMMDD")
                                }
                            >
                                <button
                                    className="border-2 rounded-sm px-2"
                                    style={{
                                        background: "rgb(var(--background))",
                                    }}
                                >
                                    {"->"}
                                </button>
                            </Link>
                        )}
                    </div>
                </div>
            </div>

            <div className="mt-8">
                <EventList list={events} />
            </div>
        </div>
    );
}

function getData(date: Dayjs) {
    return apicall<EventResponse>("get", "list", {
        date: date.startOf("day").format("YYYYMMDDHHmmss"),
    });
}

type EventResponse = {
    events: EventData[];
};

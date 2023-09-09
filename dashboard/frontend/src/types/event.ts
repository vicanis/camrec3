export type EventList = {
    count: number;
    items: Event[];
};

export type Event = {
    id: string;
    timestamp: string;
    processed: boolean;
    raw: string;
};

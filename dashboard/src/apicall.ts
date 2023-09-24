export default async function apicall(
    method: "get" | "post",
    action: string,
    args?: any
) {
    const url = new URL(process.env.BACKEND_ENDPOINT!);

    const headers: {
        [key: string]: string;
    } = {
        "X-Api-Key": process.env.BACKEND_APIKEY!,
    };

    const req: RequestInit = {};

    url.searchParams.set("action", action);

    if (args) {
        if (method === "get") {
            for (const key in args) {
                url.searchParams.set(key, args[key]);
            }
        } else if (method === "post") {
            req.body = JSON.stringify(args);
            headers["Content-Type"] = "application/json";
        }
    }

    if (method === "get" && action === "load") {
        headers["Accept"] = "video/mp4";
    }

    return fetch(url, { ...req, headers });
}

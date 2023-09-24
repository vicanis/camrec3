import apicall from "@/apicall";
import { NextRequest, NextResponse } from "next/server";

export async function GET(req: NextRequest) {
    const { searchParams: qsa } = new URL(req.url);

    const file = qsa.get("file");

    const resp = await apicall("get", "load", { file });

    return new NextResponse(resp.body, {
        headers: { "Content-Type": "video/mp4" },
        status: 200,
    });
}

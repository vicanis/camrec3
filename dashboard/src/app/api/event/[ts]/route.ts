import { NextRequest, NextResponse } from "next/server";

export async function GET(
    req: NextRequest,
    { params }: { params: { ts: string } }
) {
    return NextResponse.redirect(
        process.env.PROCESSOR_HOST + "/api/event/" + params.ts
    );
}

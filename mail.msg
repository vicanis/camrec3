Return-Path: <0102018a6f9b4384-cafc8281-9f15-439d-a2f3-d0e305e86039-000000@eu-west-1.amazonses.com>
Received: from a52-160.smtp-out.eu-west-1.amazonses.com (a52-160.smtp-out.eu-west-1.amazonses.com [54.240.52.160])
 by inbound-smtp.eu-west-1.amazonaws.com with SMTP id 86kti43e5ipu6uguhnnc1qh9cu331k40qlecpsg1
 for admin@cloud.fbuzz.pro;
 Thu, 07 Sep 2023 12:26:01 +0000 (UTC)
X-SES-Spam-Verdict: PASS
X-SES-Virus-Verdict: PASS
Received-SPF: pass (spfCheck: domain of eu-west-1.amazonses.com designates 54.240.52.160 as permitted sender) client-ip=54.240.52.160; envelope-from=0102018a6f9b4384-cafc8281-9f15-439d-a2f3-d0e305e86039-000000@eu-west-1.amazonses.com; helo=a52-160.smtp-out.eu-west-1.amazonses.com;
Authentication-Results: amazonses.com;
 spf=pass (spfCheck: domain of eu-west-1.amazonses.com designates 54.240.52.160 as permitted sender) client-ip=54.240.52.160; envelope-from=0102018a6f9b4384-cafc8281-9f15-439d-a2f3-d0e305e86039-000000@eu-west-1.amazonses.com; helo=a52-160.smtp-out.eu-west-1.amazonses.com;
 dkim=pass header.i=@amazonses.com;
 dkim=pass header.i=@hicloudcam.com;
 dmarc=none header.from=hicloudcam.com;
X-SES-RECEIPT: AEFBQUFBQUFBQUFIRDN4R2ZEa055Y1lrQVRRZkdXNS9uOHA0Y2gva2hHcWdwVGNEK0JPS0M1YzRER1VrNEZDS0Erem5tODNKS0F2bDNzOHVZOVJEUkRiRk82ZUlRTkVaSVRiRW9GdkludkZoRFpxZEYxelVicXBGMGw1U2V4bGY5SVFZNkRWS2xRNHdvYTdvME04eTQzZWNrenhBS3J4U1VLUGpUVGdrZ05CM0R0YWZhV0hIeDU3QzJXUW0yNWxmZXdyb01VdDNib0t5bld1Z3ZrZ1dwUG9aWGs5SnAyYzR6MkdBamZ1c0tKY3c4amlQVkN3ejBVM21ONThoNXFPQ3NoTjlEUml2bXpaRk8zOXgxZFF4c3ZqSW9oSzN4cm5UZzh3OEl0SHVLUGJQdHg3d2ZmV0dxNEoxSk5SejFpbXJPbDREd1ljTkd5ckE2ZEFGeWxsckgzdVg2ZURmdlJtTk5zdDRscGQ0V1JDN3BGU1hXTVhzT2s1UVZ2M0dMUkhnell2Y1VMYTN4N1BNPQ==
X-SES-DKIM-SIGNATURE: a=rsa-sha256; q=dns/txt; b=env5qp7o64Tdiaftz7kN0ngzK4pyPchNoN+tlnOmksuww3mwTk2WR9uwi+BUAS/IBsMvwU8idPwgJ3RrnzIXTRQOTu4Mt7z/KAdMrthfi9ApdHLswAyWMOJUA0HzCoLxobC7B7d8xY/xOBU8ILZ9wxk1SCc9G7qO0X5WPWFPfMc=; c=relaxed/simple; s=shh3fegwg5fppqsuzphvschd53n6ihuv; d=amazonses.com; t=1694089562; v=1; bh=I/4+RauNz+Ceax+aVbPkCnu52DQFbFSk2KoWAiX2uBA=; h=From:To:Cc:Bcc:Subject:Date:Message-ID:MIME-Version:Content-Type:X-SES-RECEIPT;
DKIM-Signature: v=1; a=rsa-sha256; q=dns/txt; c=relaxed/simple;
	s=rze7kh72auqwlyqrbz2xulpdp5cdxzsr; d=hicloudcam.com; t=1694089561;
	h=Date:From:To:Message-ID:Subject:MIME-Version:Content-Type:Content-Transfer-Encoding;
	bh=I/4+RauNz+Ceax+aVbPkCnu52DQFbFSk2KoWAiX2uBA=;
	b=JovAIxK7jTVPh+kPC2/mOs/XCNXXJA6pnlAE68Vz7OXVqaaGgfTKwLxuD93pZR2t
	lo9LboCAJAXVdcR5j1neoLiRaXFLM/pnXoR1fpYf/zK3+2vc8/Pcv6glzvvYBXitO5C
	JoCZIKFP2W0rTUWs2ymF9BszurvJdD+BunB/cFv8=
DKIM-Signature: v=1; a=rsa-sha256; q=dns/txt; c=relaxed/simple;
	s=shh3fegwg5fppqsuzphvschd53n6ihuv; d=amazonses.com; t=1694089561;
	h=Date:From:To:Message-ID:Subject:MIME-Version:Content-Type:Content-Transfer-Encoding:Feedback-ID;
	bh=I/4+RauNz+Ceax+aVbPkCnu52DQFbFSk2KoWAiX2uBA=;
	b=pmx5g2sqcxH4gWp5cAxuUXI6//Y52ki8MRtSLqoplT+2wxftRs+tgc2dbxuZYgGj
	mLSzryD8XpJFLiOsKROwXZL+JGmImJhG9+jpd8W0QEajk/nFMCKjDCM/hBUDSnT4K30
	VMTq3pDrzybRmFo2FmxVIg7rnPnI45RHrt66kBTY=
Date: Thu, 7 Sep 2023 12:26:00 +0000
From: no_reply@hicloudcam.com
To: admin@cloud.fbuzz.pro
Message-ID: <0102018a6f9b4384-cafc8281-9f15-439d-a2f3-d0e305e86039-000000@eu-west-1.amazonses.com>
Subject: =?utf-8?Q?C3WN(K49112334)_?=
 =?utf-8?Q?=D0=A1=D0=B8=D0=B3=D0=BD=D0=B0=D0=BB_=D0=BE?=
 =?utf-8?Q?=D0=B1=D0=BD=D0=B0=D1=80=D1=83=D0=B6=D0=B5=D0=BD?=
 =?utf-8?Q?=D0=B8=D1=8F_=D0=B4=D0=B2=D0=B8=D0=B6=D0=B5?=
 =?utf-8?Q?=D0=BD=D0=B8=D1=8F_From?= =?utf-8?Q?_C3WN(K49112334)?=
MIME-Version: 1.0
Content-Type: text/html;charset=utf-8
Content-Transfer-Encoding: quoted-printable
Feedback-ID: 1.eu-west-1.SfE/PwkzjUCxSwdNrIGKKphN6mrlz288Hgghuglgodk=:AmazonSES
X-SES-Outgoing: 2023.09.07-54.240.52.160

<div><img style=3D"width:100%;height:150px;max-width:375px;"src=3D'https://=
pmsas1.ezvizlife.com:8444/image/customVoice/1751030810b24538b62f364493e8bd4=
e'/></div><div><p>C3WN(K49112334) =D0=A1=D0=B8=D0=B3=D0=BD=D0=B0=D0=BB =D0=
=BE=D0=B1=D0=BD=D0=B0=D1=80=D1=83=D0=B6=D0=B5=D0=BD=D0=B8=D1=8F =D0=B4=D0=
=B2=D0=B8=D0=B6=D0=B5=D0=BD=D0=B8=D1=8F</p><p>C3WN(K49112334)</p></p>2023-0=
9-07 16:25:44</p></div></br><div style=3D"color:grey"><p>You can view more =
via EZVIZ APP.</p><p>Getting too many email alerts?</p><p>You can disable t=
he alerts by going to Menu>General Settings>Email Alerts.</p></div>

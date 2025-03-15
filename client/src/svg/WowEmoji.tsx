import { SVGProps } from "react";

export function WowEmoji(props: SVGProps<SVGSVGElement>) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      className="size-9"
    >
      <g fill="none">
        <circle
          cx="12"
          cy="12"
          r="10"
          stroke="currentColor"
          strokeLinecap="round"
          strokeWidth="2"
        ></circle>
        <circle
          cx="9"
          cy="9"
          r="1.25"
          fill="currentColor"
          stroke="currentColor"
          strokeLinecap="round"
          strokeWidth=".5"
        ></circle>
        <circle
          cx="15"
          cy="9"
          r="1.25"
          fill="currentColor"
          stroke="currentColor"
          strokeLinecap="round"
          strokeWidth=".5"
        ></circle>
        <path
          fill="currentColor"
          d="M15 15.5c0 1.38-1.343 2.5-3 2.5s-3-1.12-3-2.5s1.343-2.5 3-2.5s3 1.12 3 2.5"
        ></path>
      </g>
    </svg>
  )
}

import React from "react"

export function useIsMobile(breakpoint = 768) {
    const [isMobile, setIsMobile] = React.useState<boolean>(() =>
        typeof window !== "undefined" ? window.innerWidth < breakpoint : false
    )

    React.useEffect(() => {
        if (typeof window === "undefined") return
        const mq = window.matchMedia(`(max-width: ${breakpoint}px)`)
        const handler = () => setIsMobile(mq.matches)
        handler()
        mq.addEventListener?.("change", handler)
        return () => mq.removeEventListener?.("change", handler)
    }, [breakpoint])

    return isMobile
}
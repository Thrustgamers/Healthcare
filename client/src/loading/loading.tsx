import type React from "react";
import { Loader } from '@mantine/core';

type props = { color: "dark" | "light" } 

const Loading: React.FC<props> = (data) => {

    return (
        <>
            <Loader color={data.color === "dark" ? 'gray' : 'rgba(255, 255, 255, 1)'} size="xl" />;
        </>
    )
}
export default Loading
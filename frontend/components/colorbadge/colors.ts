import { xxHash32 } from 'js-xxhash';

// random-ish seed - non-cryptographic so it's fine
const seed = 14049382849;

function hashString(str: string, max: number): number {
    let hashNum = xxHash32(str, seed);
    return Number(hashNum.toString(10)) % max;
}

export function getColorFromString(str: string): string {
    const colors = [
        "#ff0000",
        "#ff2e2e",
        "#ff5d5d",
        "#ff8b8b",
        "#ffb9b9",
        "#0500ff",
        "#322eff",
        "#5f5cff",
        "#8c8aff",
        "#b9b8ff",
        "#ff6b00",
        "#ff862e",
        "#ffa15c",
        "#ffbb8a",
        "#ffd6b8",
    ];

    return colors[hashString(str, colors.length)];
}


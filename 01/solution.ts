const text = await Deno.readTextFile("data.txt");
const delimeter = "   ";
const left: number[] = [];
const right: number[] = [];

text.split("\r\n").forEach((str: string) => {
    const [l, r] = str.split(delimeter);

    left.push(parseInt(l));
    right.push(parseInt(r));
});

left.sort();
right.sort();

let totalDistance = 0;

// We know left and right are the same size so this works
left.forEach((l, i) => {
    const r = right[i];

    if (l > r) {
        totalDistance += l - r;
        return;
    }

    totalDistance += r - l;
});

console.log(totalDistance);

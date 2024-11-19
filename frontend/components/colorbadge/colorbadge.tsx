import { Badge } from "@/components/ui/badge";
import { cn } from "@/lib/utils";
import ColorHash from "color-hash";
import { getColorFromString } from "./colors";

var colorHash = new ColorHash();

export default function ColorBadge({ text }: { text: string }) {
  let hex = getColorFromString(text);
  let rgb = hexToRgb(hex);

  let isDark = false;
  if (rgb && rgb.r * 0.299 + rgb.g * 0.587 + rgb.b * 0.114 > 186) {
    isDark = true;
  }

  const textExplicitClass = isDark ? "text-black" : "text-white";

  return (
    <Badge
      key={text}
      style={{
        backgroundColor: getColorFromString(text),
      }}
      className={cn("text-nowrap", textExplicitClass)}
    >
      {text}
    </Badge>
  );
}

function hexToRgb(hex: string) {
  var result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex);
  return result
    ? {
        r: parseInt(result[1], 16),
        g: parseInt(result[2], 16),
        b: parseInt(result[3], 16),
      }
    : null;
}

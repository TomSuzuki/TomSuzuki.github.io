import { scroll_to_id } from "./common.js";

// works_page ...works page switching and scroll works top
export function works_page(now, next) {
    document.getElementById("works_page_" + now).style.display = "none";
    document.getElementById("works_page_" + next).style.display = "block";
    scroll_to_id("works", -16);
}
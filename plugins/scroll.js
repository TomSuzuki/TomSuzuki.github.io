// scroll_to_id ...scroll(ID、correction（px）、time)
var isScrolling = false;
const scroll_to_id = (id, correction = 0, ms = 800) => {
    // kill scroll
    if (isScrolling) return;
    isScrolling = true;

    // init
    const from = window.pageYOffset;
    const distance = document.getElementById(id).getBoundingClientRect().top + correction;
    const fps = 60;

    // scroll
    let count = 0;
    let interval = setInterval(() => {
        count++;
        let y = from + distance * easeOutExpo(count / (fps * ms / 1000));
        scrollTo(0, y);
        if (count == fps * ms / 1000) killScroll();
    }, 1000 / fps);

    // for kill
    window.addEventListener('wheel', () => {
        killScroll();
    });

    return;

    // kill
    function killScroll() {
        clearInterval(interval);
        isScrolling = false;
    }

    // easeing
    function easeOutExpo(x) {
        return x === 1 ? 1 : 1 - Math.pow(2, -10 * x);
    }
}

export default ({ }, inject) => {
    inject('scroll_to_id', scroll_to_id);
}
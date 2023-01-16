import { createApp } from "vue";
import { createPinia } from "pinia";

import App from "./App.vue";

import router from "./router";

import "./assets/main.css";

const app = createApp(App);

app.use(createPinia());
app.use(router);

app.mount("#app");

/*
    TAILWIND NOTLAR

    1)
    modifiers:
    Pseudo-classes, like :hover, :focus, :first-child, and :required
    Pseudo-elements, like ::before, ::after, ::placeholder, and ::selection
    Media queries, like responsive breakpoints, dark mode, and prefers-reduced-motion
    Attribute selectors, like [dir="rtl"] and [open]
    These modifiers can even be stacked to target more specific situations


    2)

 */

/*
    talwind
    todo: oku ve dene


    inset
    skew
    aria
    overflow
    flex-shrink, flex-grow
    letter spacing
    first-line specifier
    içinde img bulunan box'a overflow:hidden vermek
    text-transform
    text-slate rengi

 */
/*
    what is difference between flex basis and width?
    hala neden positin:relative  top:x bottom:x left:x right:x veriliyor?
    şu max-width neden hiç aklıma gelimiyor kullanmak

    todo: sor:sanırsam tailwind'de ya da herhangi bir şeyde her bir screen size için container farklı bir max-width kullanıyor. ama biz hep sabit bir max-w kullanıyoruz. neden ve hangisi doğru?
 */

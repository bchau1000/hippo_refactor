# Implementation Guidelines
- Import and export general components into `index.ts` to make importing to other components cleaner. For example:
    ```ts
    // With index.ts
    import { Component1, Component2 } from "General/index";
    ```
    ```ts
    // Without index.ts
    import Component1 from "General/Component1.svelte";
    import Component2 from "General/Component2.svelte";
    ```
- General components should all have a `style` prop that allows the caller to override/add styles passed to classes. For example:
    ```jsx
    // SimpleInput Component
    <input class={`${style} font-black`}/>

    // Caller
    <SimpleInput style={"text-lg"}/>

    // Resulting DOM
    <input class="text-lg font-black">
    ```
- General components should not have to manage their own state -- should be the job of the caller.
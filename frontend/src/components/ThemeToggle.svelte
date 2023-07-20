<script lang="ts">
    import Button from "@smui/button";

    let isDarkTheme = true;

    function toggleTheme() {
        isDarkTheme = !isDarkTheme;

        // Update the stylesheet based on the theme
        const themeLink = document.getElementById("theme-link");
        if (themeLink) {
            let newHref = isDarkTheme ? "src/smui-dark.css" : "src/smui.css";
            themeLink.setAttribute("href", newHref);
        }
    }
</script>

<Button on:click={toggleTheme}>Toggle Theme</Button>
<svelte:head>
    <link id="theme-link" rel="stylesheet" href="src/smui-dark.css" />
</svelte:head>

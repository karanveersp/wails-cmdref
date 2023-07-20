<script lang="ts">
    import FuzzyFinder from "./components/FuzzyFinder.svelte";
    import ThemeToggle from "./components/ThemeToggle.svelte";
    import ProgramDrawer from "./components/ProgramDrawer.svelte";
    import { GetCommands } from "../wailsjs/go/main/App";
    import { onMount } from "svelte";

    let commands: Command[] = [];
    onMount(() => {
        GetCommands()
            .then((data) => {
                const jsonData = JSON.parse(data);
                commands = jsonData;
                console.log(commands);
            })
            .catch((err) => {
                console.error(err);
            });
    });
    // TODO:
    // - Add Fuzzy finder logic
    // - Add Tooltips to buttons
    // - Add New Command Modal Dialog Form and connect to Go function to create and save new command.
    // - Add Edit Command Modal Dialog Form and connect to Go function to save the command.
    // - Add Delete Command Modal Confirmation Dialog and connect to Go function to remove command from file.
    // - Add Import button at the top to import from .json file.
</script>

<main>
    <h2>CMDREF</h2>

    <div class="theme-toggle">
        <ThemeToggle />
    </div>

    <FuzzyFinder />

    <div class="program-drawer">
        <ProgramDrawer {commands} />
    </div>
</main>

<style>
    .theme-toggle {
        float: right;
    }
    .program-drawer {
        padding: 2em;
        height: 100%;
    }
</style>

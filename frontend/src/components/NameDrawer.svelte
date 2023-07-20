<script lang="ts">
    import Drawer, { AppContent, Content, Header, Title } from "@smui/drawer";
    import List, { Item, Text } from "@smui/list";
    import CommandView from "./CommandView.svelte";

    export let commands: Command[] = [];
    export let platform: string = "";

    let command: Command;

    $: filtered =
        platform === ""
            ? []
            : commands.filter((cmd) => cmd.platform === platform);
</script>

<div class="drawer-container">
    <Drawer>
        <Content>
            <Header>
                <Title>Command</Title>
            </Header>
            {#if platform === ""}
                <span>No platform chosen</span>
            {:else}
                <List>
                    {#each filtered as cmd}
                        <Item on:click={() => (command = cmd)}>{cmd.name}</Item>
                    {/each}
                </List>
            {/if}
        </Content>
    </Drawer>

    <AppContent>
        <CommandView {command} />
    </AppContent>
</div>

<style>
    span {
        padding-left: 1em;
    }
    .drawer-container {
        position: relative;
        display: flex;

        z-index: 0;
        height: 500px;
        width: 100%;
    }
</style>

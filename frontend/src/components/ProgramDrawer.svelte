<script lang="ts">
    import Drawer, { AppContent, Content, Header, Title } from "@smui/drawer";
    import List, { Item } from "@smui/list";
    import NameDrawer from "./NameDrawer.svelte";

    export let commands: Command[] = [];

    $: platforms = Array.from(new Set(commands.map((cmd) => cmd.platform)));
    let clicked;
</script>

<div class="drawer-container">
    <Drawer>
        <Content>
            <Header>
                <Title>Program</Title>
            </Header>
            {#if platforms.length === 0}
                No commands found.
            {:else}
                <List>
                    {#each platforms as platform}
                        <Item on:click={() => (clicked = platform)}>
                            {platform}
                        </Item>
                    {/each}
                </List>
            {/if}
        </Content>
    </Drawer>

    <AppContent>
        <NameDrawer {commands} platform={clicked} />
    </AppContent>
</div>

<style>
    .drawer-container {
        position: relative;
        display: flex;

        border: 1px solid
            var(--mdc-theme-text-hint-on-background, rgba(0, 0, 0, 0.1));
        overflow: hidden;
        height: 500px;
        background-color: var(--mdc-theme-surface);
    }
</style>

<script lang="ts">
    import Paper, { Content } from "@smui/paper";
    import { Icon } from "@smui/common";
    import IconButton from "@smui/icon-button";

    export let command: Command;
</script>

<div class="cmdView">
    {#if command === undefined}
        No command chosen.
    {:else}
        <div class="buttons">
            <IconButton>
                <Icon class="material-icons">content_copy</Icon>
            </IconButton>
            <IconButton>
                <Icon class="material-icons">edit</Icon>
            </IconButton>
            <IconButton>
                <Icon class="material-icons">delete</Icon>
            </IconButton>
            <br />
            <br />
            <span class="cmd">{command.command}</span>
            <br />
            <br />
            <Paper color="secondary">
                <Content>
                    {command.description}
                </Content>
            </Paper>
        </div>
    {/if}
</div>

<style>
    .buttons {
        padding-top: 1em;
        padding-left: 1em;
    }
    .cmdView {
        padding: 1em;
    }
    .cmd {
        color: var(--mdc-theme-primary);
        cursor: pointer;
    }
</style>

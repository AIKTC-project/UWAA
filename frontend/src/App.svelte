<script lang="ts">
	import Welcome from "./components/Welcome.svelte";
</script>

<div class="cont">
	<Welcome />
</div>

<style lang="scss">
	.cont {
		display: flex;
		align-items: center;
		justify-content: center;
		flex-direction: column;
		gap: 1rem;
		height: 100%;
	}
</style>

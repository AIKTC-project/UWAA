<script lang="ts">
	let name = "";
</script>

<div class="head">Welcome {name}</div>
<input
	class="inp"
	type="text"
	placeholder="Enter your name"
	bind:value={name}
/>

<style lang="scss">
	.head {
		color: red;
		font-size: 2rem;
	}

	.inp {
		padding: 0.6rem;
		font-size: 1.2rem;
	}
</style>

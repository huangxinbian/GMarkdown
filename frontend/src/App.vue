<template>
	<div id="app">
		<Home/>
	</div>
</template>

<script>
	import Home from './components/Home.vue'

	export default {
		name: 'app',
		components: {
			Home
		}
	}
</script>

<style>
	html, body, #app {
		width: 100%;
		height: 100%;
		padding: 0;
		margin: 0;
		/* overflow: hidden; */
	}
	
	.not-select {
		-moz-user-select: none;
		/*火狐*/
		-webkit-user-select: none;
		/*webkit浏览器*/
		-ms-user-select: none;
		/*IE10*/
		-khtml-user-select: none;
		/*早期浏览器*/
		user-select: none;
	}
</style>
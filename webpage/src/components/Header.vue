<template>
	<mu-appbar class="header" title="葫芦兄弟大战多线程白骨精">
		<mu-icon-button v-on:click="switchSidebar" icon="menu" slot="left"/>
	</mu-appbar>
</template>

<script>
	export default {
		name: 'GOHeader',
		props: {},
		data () {
			return {}
		},
		methods: {
			// 触发改变侧边栏显示状态
			switchSidebar: function () {
				this.$emit('switch')
			}
		},
		mounted: function () {
		}
	}
</script>

<style lang="stylus">
	.header
		top 0
		position fixed
		text-align left
</style>

<template>
	<div id="app">
		<GOHeader v-on:switch="handleSwitchSidebar" />
		<Sidebar v-bind:isHidden="hideSidebar"></Sidebar>
		<div class="main-content" v-bind:class="{ hideSidebar: hideSidebar }" >
			<div class="container">
				<router-view/>
			</div>
		</div>
		<AlertDialog ref="alertDialog"/>
	</div>
</template>

<script>
	import AlertDialog from './components/AlertDialog.vue'
	import Sidebar from './components/SideBar.vue'
	import GOHeader from './components/Header.vue'
	export default {
		name: 'app',
		components: {
			AlertDialog,
			Sidebar,
			GOHeader
		},
		data () {
			return {
				hideSidebar: false
			}
		},
		methods: {
			handleSwitchSidebar () {
				this.hideSidebar = !this.hideSidebar
			}
		},
		mounted: function () {
			this.mountDialog(this.$refs['alertDialog'])
		}
	}
</script>

<style lang="stylus">
	THEME_COLOR = #2196f3

	body
		margin 0
		padding 0
	#app
		font-family 'PINGFANG SC', 'Avenir', Helvetica, Arial, sans-serif
		-webkit-font-smoothing antialiased
		-moz-osx-font-smoothing grayscale
		text-align center
		color #2c3e50

	.main-content
		margin-left 200px
		margin-top 60px
		transition all 0.3s ease
		padding 16px
		overflow-y auto
		&.hideSidebar
			margin-left 0
			.container
				margin 0 auto
		.container
			transition all .6s ease
			margin 16px
			/*width 1280px*/

	.section
		text-align left
		.title
			border-bottom 2px solid THEME_COLOR
		.introduction
			.high-light
				color red
		.panel
			display flex
			flex-flow row nowrap
			justify-content space-between
			align-items flex-start
			.introduction
				margin 0 16px 0 0
			.source-controls
				display flex
				flex-flow row
		.source
			display flex
			flex-flow row
			justify-content center
			align-items flex-start
			.thread
				padding 16px
				margin-right 16px
				flex 1
				box-shadow 0 0 4px 1px rgba(0, 0, 0, 0.16)
				&:last-child
					margin-right 0
				.thread-header
					font-size 20px
			.code
				margin 32px 0 16px
				.instruction
					font-family "Fira Code"
					font-size 16px
					.indent
						white-space pre
					&.current
						background #f3ff74
					.block
						.comment
							font-family "PINGFANG SC"
						.keyword
							color perple
						.critical
							color THEME_COLOR
		.context
			margin 32px 0 0
			.variable

				font-size 18px
		.subtitle
			margin 16px 0
			border-bottom 1px solid THEME_COLOR




</style>

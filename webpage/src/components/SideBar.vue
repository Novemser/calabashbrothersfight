<template>
	<mu-paper v-bind:class="{ hidden: isHidden }" class="sidebar"  :zDepth="1">
		<mu-list>
			<router-link v-for="(navigator, index) in navigators" :key="index" v-bind:to="navigator['link']">
				<mu-list-item>
					<div class="item">
						<mu-icon value="settings" :size="20"></mu-icon>
						<span>{{ navigator['name'] }}</span>
					</div>
				</mu-list-item>
			</router-link>
		</mu-list>
	</mu-paper>
</template>

<script>
	export default {
		name: 'sidebar',
		props: {
			isHidden: {
				type: Boolean,
				require: true
			}
		},
		data () {
			return {
				navigators: [
					{
						name: 'introduction',
						link: '/'
					},
					{
						name: 'section-beginning',
						link: '/section/beginning'
					}
				]
			}
		}
	}
</script>

<style lang="stylus">
	$SIDE_BAR_WIDTH = 200px

	.sidebar
		transition all 0.3s ease
		width $SIDE_BAR_WIDTH
		height 100%
		position fixed

		.menu
			margin-top 24px

		.item
			color #000
			display flex
			align-items center

			span
				margin-left 8px


		&.hidden
			margin-left - $SIDE_BAR_WIDTH

</style>

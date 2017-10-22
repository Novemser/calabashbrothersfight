<template>
	<mu-paper v-bind:class="{ hidden: isHidden }" class="sidebar"  :zDepth="1">
		<mu-list>
			<router-link v-for="(navigator, index) in navigators" :key="index" v-bind:to="navigator['link']">
				<mu-list-item>
					<div class="item">
						<mu-icon value="star" :size="20"></mu-icon>
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
						name: '首页',
						link: '/'
					},
					{
						name: '你渴望力量吗',
						link: '/level/0'
					},
					{
						name: '魔鬼般的赋值语句',
						link: '/level/1'
					},
					{
						name: '装甲再厚也能打破',
						link: '/level/2'
					},
					{
						name: '武功再高也怕菜刀',
						link: '/level/3'
					},
					{
						name: '穿越时空去见你',
						link: '/level/4'
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

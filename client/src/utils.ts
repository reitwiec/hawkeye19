export type Attributes<
	C extends keyof JSX.IntrinsicElements
> = React.PropsWithoutRef<JSX.IntrinsicElements[C]>;


export enum RegionKey {
	region0 = 'region0',
	region1 = 'region1',
	region2 = 'region2',
	region3 = 'region3',
	region4 = 'region4',
	region5 = 'region5',
}

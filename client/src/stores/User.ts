import { action, observable } from 'mobx';

class UserStore {
	@observable username = '';
	@observable email = '';
	@observable region1 = '';
	@observable region2 = '';
	@observable region3 = '';
	@observable region4 = '';
	@observable region5 = '';

	@action setCurrentUser = user => {
		this.username = user.username;
		this.email = user.email;
		this.region1 = user.region1;
		this.region2 = user.region2;
		this.region3 = user.region3;
		this.region4 = user.region4;
		this.region5 = user.region5;
	};

	@action clear = () => {
		this.username = '';
		this.email = '';
		this.region1 = '';
		this.region2 = '';
		this.region3 = '';
		this.region4 = '';
		this.region5 = '';
	};
}

export default new UserStore();
export { UserStore };

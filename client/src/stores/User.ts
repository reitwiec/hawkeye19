import { action, observable } from 'mobx';

class UserStore {
	@observable username = '';
	@observable email = '';
	@observable region0 = 0;
	@observable region1 = 0;
	@observable region2 = 0;
	@observable region3 = 0;
	@observable region4 = 0;
	@observable region5 = 0;
	@observable isVerified = 0;

	@action setCurrentUser = user => {
		this.username = user.username;
		this.email = user.email;
		this.region0 = user.region0;
		this.region1 = user.region1;
		this.region2 = user.region2;
		this.region3 = user.region3;
		this.region4 = user.region4;
		this.region5 = user.region5;
		this.isVerified = user.isVerified;
	};

	@action clear = () => {
		this.username = '';
		this.email = '';
		this.region0 = 0;
		this.region1 = 0;
		this.region2 = 0;
		this.region3 = 0;
		this.region4 = 0;
		this.region5 = 0;
		this.isVerified = 0;
	};
}

export default new UserStore();
export { UserStore };

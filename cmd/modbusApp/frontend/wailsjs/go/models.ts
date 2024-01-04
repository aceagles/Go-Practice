export namespace main {
	
	export class groupDefinition {
	    startReg: number;
	    numReg: number;
	
	    static createFrom(source: any = {}) {
	        return new groupDefinition(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.startReg = source["startReg"];
	        this.numReg = source["numReg"];
	    }
	}

}


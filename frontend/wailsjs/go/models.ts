export namespace models {
	
	export class SitemapConfig {
	    Dsturl: string;
	    Savepath: string;
	    Filename: string;
	    Concurrency: number;
	    Crawltimeout: number;
	    Requesttimeout: number;
	
	    static createFrom(source: any = {}) {
	        return new SitemapConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Dsturl = source["Dsturl"];
	        this.Savepath = source["Savepath"];
	        this.Filename = source["Filename"];
	        this.Concurrency = source["Concurrency"];
	        this.Crawltimeout = source["Crawltimeout"];
	        this.Requesttimeout = source["Requesttimeout"];
	    }
	}
	export class Config {
	    Status: number;
	    Port: number;
	    AutoProxy: boolean;
	    SaveLogFile: boolean;
	    Filter: boolean;
	    FilterHost: string;
	    Sitemap: SitemapConfig;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Status = source["Status"];
	        this.Port = source["Port"];
	        this.AutoProxy = source["AutoProxy"];
	        this.SaveLogFile = source["SaveLogFile"];
	        this.Filter = source["Filter"];
	        this.FilterHost = source["FilterHost"];
	        this.Sitemap = this.convertValues(source["Sitemap"], SitemapConfig);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}


package main

import (
	"reflect"
	"testing"
)

func Test_day3_compareGroups(t *testing.T) {
	type args struct {
		groups [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test case 1 - pass",
			args: args{
				groups: [][]string{
					{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"},
					{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
				},
			},
			want:    "rZ",
			wantErr: false,
		},
		{
			name: "test case 2 - fail",
			args: args{
				groups: [][]string{
					{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"},
					{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
				},
			},
			want:    "",
			wantErr: true,
		},

		{
			name: "test case 3 - day3_input",
			args: args{
				groups: [][]string{
					{"FqdWDFppHWhmwwzdjvjTRTznjdMv", "ZBJrDVfQcfSRMLjZnjjM", "cBffPfbrbQcgQJggfVQJBPbCwlPtWFDWHFHhpmmGlGmlqmDG"},
					{"ljvvqhlvshhnrcpBZqpTcr", "zGhWzFTJvsFttddWbMRdmP", "hhGdDhfdDhmdnHwtzzRtdzbzQQQg"},
					{"vLJcLLTRsvsTZPqHGhFHGhhPhfqD", "dJszLvzvCZZsJmszCrrdFmpppMpDMQPMgmDcDgSS", "nRjRbnnjlNhblnjtVtQlWttMFPfMfPgDMpBgSBPgBS"},
					{"RnLJBfmJfmNBHlQvvbdQ", "grhgrtqgjJhhggNHqvwWqvbNlbHw", "GVTTsFFjJjVVFVGCFTJDDjhFcZmRMZfnZcncSpMSGcRPZpLp"}, {"FGpDqTTVFFprpjLVQMMGtMclcmHGtBdc", "TjhhgTLpVZhpLDZqrTqZVpBSwJwzNnWNWJvzzNSggwwN", "PHTMsmwrJMwLJvJddvdHwvcWnnWfccqGnhhfGcDqsnGc"}, {"CmlgmZlQtFtZNlVZdHLMrMvvCrrvvTTC", "PrDGBBddprmzddrSqccRgSTpqbsMRR", "hvLtfFNvvZNfGGfRgbqsRNlTSSgsbM"}, {"GrjGrpjjCsnwhsGGPwlPTPLPVttPqLVl", "HzSHHhczRlLTHqqq DbhvFSSzQcZbcFbcQjrJrMJmmZnGJmJnjn"}, {"BVtPqFMqtvQFqPqjFBMVtRZGNGhfNcfQdpfgnQgNcNgp", "HBHtPbHCLCzsLJvT", "nVHVFfggbQVmFFfhLpBpBTrfLBCB"}, {"tJqStcRMMMjPwlnvzgQWzmHzlmQl", "nNSRcDHRmHhhDZZZdBDfWJdfrJ", "LGrGVtjCPCbbQQQQLvQpbVQZzsqsBMdBqMvZMMJZqJvdBW"}, {"rQbnBrDTQcdpHttt", "qNsNpfjLpNLRNqLMtdVsddcmVzdzVh", "WJfLCfvpWpNLbwwwBwbZFvlZ"},
					{"qtDCZcfZtDjbjCQvmmSj", "ttpRzqtqRLDzpRtDdtPlhhZZGBHGPsGZPhwd", "WjtcNnMtztGFrsNjcgRHvdwHhhjRwlhwlg"}, {"qpDBpBqPTDLqGLmpTPqbDmWMNrFrrFQnFMNnMNQtMm", "dNwfsjFLQLFNBhTCTPPTBJhhTP", "VmVMgqgRVHtztmgqgzgqRzgMPWCRWTpPCCPWPThDdZPCRZJp"}, {"qbWcqrFNCJGSChvLGv", "MnslRSpSVsGgGhDDwLvl", "RjMtfnpmmmjSWSbjrNPN"},
					{"ZRsBvMvmZMGQQmFnDmQzzQ", "MlGLlBjRGjjWWGRGMlNrLLrssTPNfLcTgrVr", "vpSBBcJnWnSmcqmcChdcgf"}, {"DzPDNVNbLwPzMLRbNMDjBlvpWjrvrVJjjSsBjr", "dndGpnWdVnBPFFHTBgGH", "qCcCmjqNNJDtcJQjqJqMCQJcFTBRZRwBZBHDPRHBHZDbgBBR"}, {"NjJLgSLGGCLWvqNqNBvwFb", "htmVmtHnlZmDVhtmhblmDMHddBRvwwZQQqWPFvFwdvRWdR", "HnlnHfDsbSTbJzsz"},
					{"TzcczTtgqpgGSTlHHVHCGjCdHdnQQQ", "wDBDSDlWzwwzDqqMtrMrfsfbNvJNJB", "ThfQTdQzQbgdhdNbJFJSlbBLlmqqHBSHCRHsHm"}, {"cGWPPMtpZDwpMpJQBQQQDBBVNfdQ", "WcWDRLSzFrRFFccPmnssMrGtmMnnGm", "jvvgCCTvNqTtJqsnqwPBmspZMB"},
					{"zTrHgrFWRrWvMpPNBVZZHVfN", "hdnlltGLtGSQPVLSNBQN", "CqGGtGwlhlGdtGmbtjtmmvDRFvVFTTRDRbRgFJRbDr"}, {"zzPqvdqNzvqzfzMfzqmCzzfJsclWjSSRWglgjclWSmcjlJ", "hfPGwhhvMGCVCdddhQQZQnNrTDQnFrtn", "DDMFjwVTgVmMWgVpdqtlJnpvHHnslw"},
					{"GLcNhRGLZBhmDWTjDTWF", "FbVcQRVRBFfNFfccVfZcWddnGrrHncWtdHsZ", "hwGwjgTSGrssZHHTdn"}, {"rfMCJPBMMCrSCSBGZZqRlRLzqhqh", "LDvdNmbgHjHgnmnvnHjgDjqlZhZzszhqzWsRWRRNcRNZ", "dmbvHDdnjDjVmjTmHjJVrfprPCCVtLSMrSpJ"}, {"FFWFWTzWSWtFgPHgRPWTzggpJcvvSCGSGJnrvhhrrJhlGlGn", "pztFjzRTqWzgHqHWtPtPFgmsNfqbDdwDNVwbmfwbdNsb", "zCHvDWwvCwgpNRCWWHttCwvNPVTqrRrVbbsnbqQPbVsbPrqG"}, {"ZfJBmLjfJZMcdZmJffGMtHtwNCvWwwwztMzg", "rgFgllfdpFlTHfTnfnNPNtPBBVtpzmVVPmmP", "bLSSbGhGWSWmLzztcQPCQC"}, {"GSwgSdwfvdfvwgGwBLdJbjjpmFjSTRpqHmRrjptrqt", "NDMVMCsWQVCsQFjFTHjVLqjFrp", "PzWDDCNNlCWfbhZZLfBGwl"}, {"GGLfGtbfMNbbLGtzjBNnsMjjZCwDvQZHrHQvDHHHHHHrrNHq", "BthbnfPGMstmJgRPpFRTLL", "VRcdVRPTgVTLVMwmggJBwblJFlmb"}, {"pstjcqrprsHrfDpnrCnHCzRhSMhMMMRVPLMhddfhzT", "QNDQThccDghdcLLgVsrVLVlMGM", "BbbWpppFFpnfnFbBBPRMTJMRsGlJGlJVLf"}, {"lwCLwLjzLhLHCvwjGCZJbQSHdBQdDdbDtdSQ", "pnzcTVsszpncgFdJpbtDBDttDSJS", "zTPcRrfzgzCCvGRhLLqW"}, {"jfpNqSrpCqNfNSpjCqSqshNFRQJcJWRRFssgWRHWWFWHQJ", "mrnNNgNfMZwmDtZw", "ZrrFPQsQPRLcvPJvhg"}, {"tTpTGfnttqwnqQJzFdzfSdzSzr", "qVVZqfVNdnBZMNzNnPzfMqbzJvFSjSllvjBwrvrvFrlsrjJs", "LhHtDTmWmWmGDhGLWHghHLCwrwRjjsPJrSjFjFrFClvrww"}, {"WGpVMtGZplgHVWMtZpZFHJjndvFdjddJhnjLHF", "DTrRcSSccfzcCPDCTnvJvjbrrhjNvqjqJL", "zczwPDTmfZZtgLwWlp"}, {"GGwQRZNCRgDgsDDbSbSgfGFFrHfVHVVWrHWWcVFF", "RZRPgbNSJDsTsplPLsvd", "GrGNGhpnPFFBfCQCMwrVlwTC"}, {"RjWDmDdDbjcSHmRRLRRHjdDBsFFbBngBzhqbBhpqBPgFFg", "TfdNjJjmShGcWvQNQqcNFN", "gsRZRMqbDpsHHnZnngMZFPlzQRPWvRtPwtRWzlPv"}, {"BdSRjHScGMVjGdcScwLgqQqphNqNDqBBQW", "zzQClvtttrwqrwgCwp", "PQJTzvJJTtJQlvQftmfdmdmGRcSdcMSGdS"},
					{"FRSHzMJdrfRnpmpWmp", "TzJTMNNrHLJfTJccvjqwwgGLZLLc", "dCpjsGvcsLvszTrRRlRrDJ"}, {"SBbphFNtLvvSfLfc", "zTFnnZzqrjFVnZTrtwMDptbpMwMMBDzb", "JGPCjWGgJjPWGJjNchmbBRpNNDsRsNBbbRBppD"}, {"WWbrmZjbmjpbWSmcWHSbLddwvDggLFDhFrRlFFDw", "qPMVzTPQVfzvVzBQTMtRswwlDdhDghhsfLwdhs", "MMMPQBPJPzCVzvzQVtBJJMzCbmSZNWbSccHHmmZpWSZGZS"}, {"rfCqQQrfvCQNRqCnCthFVSGSGtVS", "TrzLNvWfQvrWQZNWsnHsssBBcZBMsDdZ", "CtJCddDHDDPGHCdNVLPBdLDbbGFbwZSsSSZrfFFrGbGlsb"}, {"ghRRRhQgWvmPWtHHtLdMtH", "qNgQgNgNQFVbqVQDMRZMDRBHHJHRFh", "jnTTPzWZCzrWzGRHMrhmBvJBSBBM"}, {"DJgJDgFqCGlhFDGDCWhqCwRfpbcpbsgsVscRpwbwws", "SmLZQmMVvLbsbssNQsQN", "tHZMLZZZMVtmnMHWDHjWlGhFDqCl"}, {"wwHCHLtLwnzNHVTZZV", "rJLtZtLZcCrvJwcbrGLvrcDQDhPldhGdmmPRQhhRBlPM", "zShzVhbqlbpbRNRscBNwCc"}, {"FFFZMWWZZJMmJJMFFpztlPtSllLVphttGZ", "gWhGwmwQGhVwGzBMnDFmdmDDLbCnLn", "tRRHNHHlPHPfPltllNNNRsVPFrTrTbrMMTTDMdLFCrFdPP"}, {"jDDRDVqNsRMMVFjFbtzpBlpllCBlhSLHSStH", "JJWZZTwWcmZCzQShHhzhhm", "wTrgvTwccvdcfvJvWJrvJTNzNMfGNNGjFRsjNDbfVzDs"}, {"NNCqwNSrsDqNSSgLgffDCNpBdRvvdpmHRQvBdBqvdzBB", "CgbffsDJsSsNgbgJrlcGhGtjnWcJPncjWL", "RwwTGRjGlwWNgjgfQVNmjj"}, {"cmchnhtPqLSJJbdPLntlsDGCTWWDsDRwWWWTwS", "DDDBsPGPbwhDcDcj", "rgfNgCmHMvrrttvtfmNLgrLcQQJTJhcnjTQHcZlwlQHnwj"}, {"RwsbssJjnbJwwsGPPdDLfTDLLLWvWNDGDpWD", "zVtHqzHHVcBQTdpSgvQDSpTS", "HHhFCdrrHchmrhcmwsjZwRwMJlshJJJZ"}, {"DnVWvvpvWvzrpwRWDbTvbTjTFGlfFSfNSFHGmFGNFmnPlmFS", "wDRDZzVWVjrRVjzrQgQdbgCZLCsCQtbd", "jQSgWjQmFFvHmjHWVVpZbGlbGlfGpbfGGWpf"}, {"MTzPtcrcwzTttdBJwPvFVHjMLFHQMgsSLHvs", "BHtjmmTtmDtHZjMMdNzCzCWcWZsZdsZs", "JRRJVPLwQJrVMPJLVVwChpRRhcdzcNzhzChNdW"}, {"rsMDTrgsBNBgMgDBhfhDghrtcRWJttcmGRWLGQQLJRrqRL", "lnVVjvPbwpndvVwlVCjVwtLcsmLLqWWtttlsmGGcmm", "SsjVSCZbgTBTfNZH"},
					{"jVLTHGRHjjvPGcDrfNMbnpngVbpf", "RJRTTRZJLRJQZcGGHLhHvTdmBFWFtWzqdWWQlFzqBFWt", "HDgZHpZSDpBQdRpHHRsDBNNzTvfTQqcqNNTqTqPQvl"}, {"CWMtjCWMCCLWrWVWJwBbpgtRZSHbddZHBDDb", "hjLWPZJpZptwJghSfgHTMTgRMR", "lCcrblcnlzqDsvbfffTRSMnffjHTgj"}, {"plQtRqRlGpPPPLZtmtpttRtJjHJvddLTHrLHJJjbdHvrrN", "gcMgTBWBCTczjnvNznnbgH", "FBhSwsWFWDhQZQqTTZllsT"}, {"nzHqGNMfHqMMwJLLsNBrNcBS", "qgnHflgzHlqlZzhnzssfzbZtmVvmTQvQdTtbbDbTvv", "RtQCBbJDFhJtQtZtCbMnVnPVfGPpJVJWWmwJPG"}, {"SHcljRHzNHHjSHcrgNslcczgZtDgMDBCvQQbhDBCbFBCFCCF", "dRTFRJTRTgJzSSJmzJfN", "jLbQllLvvvrQlLQBrvQmFBGzCFtPSMmSGCtPPF"}, {"bWnDbMJMFbhZSfngpfpd", "HjcrlvjRjrjlLqTqpwQgpfSQgghZgvgf", "THNTjClLHDCCpWmbtC"},
					{"CqrJTHCvDfrfwwJHRBvRGvgQmvmlWtlR", "rHzJJJDDwJCqjHTwBLccMpFhshnSSsMz", "bqVqqrDMpLFbLpJJQDMjbpZndwsvwHHswvnvnnZslG"}, {"fTSSSBtTzhPththRrrCBzVLqMMJQjVJJCLjFjjqjJq", "pQlCJQjVvVGGnjNqFbFP", "wTqtmSmTBhstmstTmWRSdLdfPdNGnhGhdPNnNnrN"}, {"vVwCqdCDvMrlDJCqrDMrPdTFWwZNbRcbmZWQbRQZWZcQWm", "fGSfhfjHGBlhSpRmpZFZNbQWjj", "ntGlgSSGgStVqdMCDCJnqd"}, {"DVdQcsccdczbtQBWvmQBBW", "PgHPNcmcqqTphlHTCGrT", "GVLCBmdLVtlrmqGCqrTCGnjGSfwzNfPzfNNGzSSw"}, {"hcsscQQsDcWbHbmLgTbVgTmlrbgq", "VwJndPThQQjdvbrb", "BHlzFFjlZGBBlZBDFSmllfGggLNbNvggtCgNrLQLbfvbfg"}, {"QFFGpfGtwgtjwvpwpGGjjBHgqzNNllHRlRllNdBRJl", "SWVDWVrDhnWhnqVnWPrhcSldBJPsRMMdPHsRdJHRMMBP", "SWZZcmnnZnSCLZDZftFFCvbfTTqTCTQF"}, {"bbctWQDjcnwtwDDDllQzLfgsPzqgrsTPLshWfP", "bwRbQRnmDwtbTjmwRJQRQVZHpCpHVdvpMHZZMHZJvd", "CfqzMCGvGqNrCFFNwcMBbnnbbtBbDStw"}, {"dJlvLVgJTRVPWWjjggCzhNGzzGZdhqrdmqrq", "zGSHWGjzpRsWVfsNwZ", "BmPCLzPlJBBvQmLFQrwwrVtfqtqZtsvwfR"}, {"ppngVjZwNZwwVJjjnnVVJJJpLrsgmsmrbctsLcLmscDggDsL", "vHvRzRPvQPPRqPTlffRTrssbtqmmtbcMcNrDrrtD", "CHWdWzWQzdChNVZjZZwB"}, {"DCzCpJppzJJDScFBzvBGnRWGWrNRQNNWGtNGqB", "vFJSzDSszzzccSZHzDDJmvSJhVwLjfVTPZLfwLgjThLZwhVh", "FMSSNScRlSGzfqWbqqcpWBhpmW"}, {"vwvCsJZZnnwVnVHsfNjDRgSjRzgRjZjN", "CRhsghlqlvjhPslQdrMndMTWdPnTMJ", "tcDNFDpDSDwDtHrSDwDtFmtpJzWTnzQJpzQGnGWQMWnnTGBd"}, {"vHPsBzvRvhCzbwbnjHLVmVbW", "ggdlpZNdZdflWTGpVMjgVbqjVmngnVbr", "lNpfllFGpZTtclDZzSWQFCWWzChPzQhS"}, {"DDflHqNNNjCCPNbT", "LFpLzfzqHzLbLvtndJhs", "mHtbGdwzmtHZrtrHtHGwrmtcFgBFRsBFcdglRsfsdgJBNN"}, {"qDwjvqjVTDDLSPpnjqVnzbhbZbrGMmHzbMHHhh", "cMrrVsdtCdVtwqdHgLjPLFLfcfZZJB", "zGSlGGvpDWWvbSSNTggZPPTTJLZPjPzJ"}, {"WZSnCCMMdMMMSJMSVZmmFqVjqjVwVjjc", "zvzTlQQQQTNGbmQTjqBjcwwwGgBHGwqB", "mthbtmlPhTrNzNhhPLzlPzWfWPJdRCWspPWSSRsWSnJp"}, {"MrggHrMpbtTpgpDptlQRJJldVRPBsDQvvV", "grgrTzfTGSFGsGSCLc", "hjdjCfQCLdQcWMfDDQhLsLCjqvVqzzZZprBFbzVdvVnqnBbd"}, {"gPGgPJSTPJtHncnjWDhDMDhP", "ZWWqBqqmgFFDMTTWDMwwcW", "ddGPpJQPprRSCrQzJPJGjHDvssMTDZDwcHhMsHsMHS"}, {"pslQSspQrqHfgMRl", "tJBTjsTLGMBGMFMg", "tDwnhZdnLdZDwczzcPvsVvVW"}, {"dlTqNqfTjdJflCppCQQRHVwFRJrvJsFVQs", "LNTqBLPprpLLzWtSnMZBMDGm", "PBjlPvvcJlJzwqjnnjLnwm"}, {"SWppdTrprThhrGVztcrllrQJZNlBZlrNZB TTrNcjGNWDdDPDpPjDPNrbmFPfQSFmfSMCmCfSmSSR", "sshhqshzBVnzHgwCMbCwcfmRFmRMbm"},
				},
			},
			want:    "DhPJDghBJptPSBPbzPVvGLWmLLszrwLNWQzBvrhVwFzNcdVMsTMTnSpqPlmQqbVNzJWfjmrTsqc",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day3_compareGroups(tt.args.groups)
			if got != tt.want && !tt.wantErr {
				t.Errorf("day3_compareGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day3_sumElvesPriorities(t *testing.T) {
	type args struct {
		ret_day3 [][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1 - pass",
			args: args{
				ret_day3: [][]string{
					{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"},
					{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
				},
			},
			want: 70,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day3_sumElvesPriorities(tt.args.ret_day3); got != tt.want {
				t.Errorf("day3_sumElvesPriorities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_day4_calculatePairs(t *testing.T) {
	type args struct {
		ret_day4 [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test case 1 - pass",
			args: args{
				ret_day4: [][]string{
					{"2-4", "6-8"},
					{"2-3", "4-5"},
					{"5-7", "7-9"},
					{"2-8", "3-7"},
					{"6-6", "4-6"},
					{"2-6", "4-8"},
				},
			},
			want:    2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day4_calculatePairs(tt.args.ret_day4)
			if got != tt.want && !tt.wantErr {
				t.Errorf("day4_calculatePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fillNumbers(t *testing.T) {
	type args struct {
		assignment string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test case 1 - pass",
			args: args{
				assignment: "24",
			},
			want:    []string{"2", "3", "4"},
			wantErr: false,
		},
		{
			name: "test case 2 - pass",
			args: args{
				assignment: "66",
			},
			want:    []string{"1", "2", "3", "4", "5", "6"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fillNumbers(tt.args.assignment)
			if !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("fillNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

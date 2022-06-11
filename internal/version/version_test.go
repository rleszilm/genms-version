package version_test

import (
	"testing"

	"github.com/rleszilm/genms-version/internal/version"
	"github.com/rleszilm/genms-version/internal/version/versionfakes"
)

func refBool(b bool) *bool {
	return &b
}

func refString(s string) *string {
	return &s
}

func TestMajor(t *testing.T) {
	testcases := []struct {
		desc       string
		branch     string
		commit     string
		committish string
		tag        string
		opts       []*version.VersionOption
		expect     string
	}{
		{
			desc:       "default",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			expect:     "v1",
		},
		{
			desc:       "default - non master branch",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Master: refString("main")},
			},
			expect: "master",
		},
		{
			desc:       "all options",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Full: refBool(true)},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v1.2.3-master+d28fbce",
		},
		{
			desc:       "all options - docker",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Docker: refBool(true)},
				{Full: refBool(true)},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v1.2.3-master-d28fbce",
		},
		{
			desc:       "all options - non master branch",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Full: refBool(true)},
				{Master: refString("main")},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v1.2.3-master+d28fbce",
		},
		{
			desc:       "all options - non master branch - docker",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Docker: refBool(true)},
				{Full: refBool(true)},
				{Master: refString("main")},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v1.2.3-master-d28fbce",
		},
	}

	fv := &versionfakes.FakeVersioner{}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			fv.BranchesReturns([]string{tc.branch}, nil)
			fv.CommitReturns(tc.commit, nil)
			fv.CommittishReturns(tc.committish, nil)
			fv.TagReturns(tc.tag, nil)

			v, err := version.NewVersion(fv, tc.opts...)
			if err != nil {
				t.Error("could not instanitate version:", err)
			}

			actual, err := v.Major()
			if err != nil {
				t.Error("could not generate version:", err)
			}

			if actual != tc.expect {
				t.Errorf("actual(%v) generated version is not as expected(%v)", actual, tc.expect)
			}
		})
	}
}

func TestMinor(t *testing.T) {
	testcases := []struct {
		desc       string
		branch     string
		commit     string
		committish string
		tag        string
		opts       []*version.VersionOption
		expect     string
	}{
		{
			desc:       "default",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			expect:     "v1.2",
		},
		{
			desc:       "default - non master branch",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Master: refString("main")},
			},
			expect: "master",
		},
		{
			desc:       "all options",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v1.2-master+d28fbce",
		},
		{
			desc:       "all options - non master branch",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Master: refString("main")},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v1.2-master+d28fbce",
		},
	}

	fv := &versionfakes.FakeVersioner{}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			fv.BranchesReturns([]string{tc.branch}, nil)
			fv.CommitReturns(tc.commit, nil)
			fv.CommittishReturns(tc.committish, nil)
			fv.TagReturns(tc.tag, nil)

			v, err := version.NewVersion(fv, tc.opts...)
			if err != nil {
				t.Error("could not instanitate version:", err)
			}

			actual, err := v.Minor()
			if err != nil {
				t.Error("could not generate version:", err)
			}

			if actual != tc.expect {
				t.Errorf("actual(%v) generated version is not as expected(%v)", actual, tc.expect)
			}
		})
	}
}

func TestPatch(t *testing.T) {
	testcases := []struct {
		desc       string
		branch     string
		commit     string
		committish string
		tag        string
		opts       []*version.VersionOption
		expect     string
	}{
		{
			desc:       "default",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			expect:     "v1.2.3",
		},
		{
			desc:       "default - non master branch",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Master: refString("main")},
			},
			expect: "master",
		},
		{
			desc:       "all options",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v1.2.3-master+d28fbce",
		},
		{
			desc:       "all options - non master branch",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Master: refString("main")},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v1.2.3-master+d28fbce",
		},
	}

	fv := &versionfakes.FakeVersioner{}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			fv.BranchesReturns([]string{tc.branch}, nil)
			fv.CommitReturns(tc.commit, nil)
			fv.CommittishReturns(tc.committish, nil)
			fv.TagReturns(tc.tag, nil)

			v, err := version.NewVersion(fv, tc.opts...)
			if err != nil {
				t.Error("could not instanitate version:", err)
			}

			actual, err := v.Patch()
			if err != nil {
				t.Error("could not generate version:", err)
			}

			if actual != tc.expect {
				t.Errorf("actual(%v) generated version is not as expected(%v)", actual, tc.expect)
			}
		})
	}
}

func TestIncMajor(t *testing.T) {
	testcases := []struct {
		desc       string
		branch     string
		commit     string
		committish string
		tag        string
		opts       []*version.VersionOption
		expect     string
	}{
		{
			desc:       "default",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			expect:     "v2",
		},
		{
			desc:       "default - non master branch",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Master: refString("main")},
			},
			expect: "master",
		},
		{
			desc:       "all options",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Full: refBool(true)},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v2.0.0-master+d28fbce",
		},
		{
			desc:       "all options - non master branch",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Full: refBool(true)},
				{Master: refString("main")},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v2.0.0-master+d28fbce",
		},
		{
			desc:       "all options - non master branch - docker",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Docker: refBool(true)},
				{Full: refBool(true)},
				{Master: refString("main")},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v2.0.0-master-d28fbce",
		},
	}

	fv := &versionfakes.FakeVersioner{}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			fv.BranchesReturns([]string{tc.branch}, nil)
			fv.CommitReturns(tc.commit, nil)
			fv.CommittishReturns(tc.committish, nil)
			fv.TagReturns(tc.tag, nil)

			v, err := version.NewVersion(fv, tc.opts...)
			if err != nil {
				t.Error("could not instanitate version:", err)
			}

			v.IncMajor()
			actual, err := v.Major()
			if err != nil {
				t.Error("could not generate version:", err)
			}

			if actual != tc.expect {
				t.Errorf("actual(%v) generated version is not as expected(%v)", actual, tc.expect)
			}
		})
	}
}

func TestIncPatch(t *testing.T) {
	testcases := []struct {
		desc       string
		branch     string
		commit     string
		committish string
		tag        string
		opts       []*version.VersionOption
		expect     string
	}{
		{
			desc:       "default",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			expect:     "v1.3",
		},
		{
			desc:       "default - non master branch",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Master: refString("main")},
			},
			expect: "master",
		},
		{
			desc:       "all options",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Full: refBool(true)},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v1.3.0-master+d28fbce",
		},
		{
			desc:       "all options - non master branch",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Full: refBool(true)},
				{Master: refString("main")},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v1.3.0-master+d28fbce",
		},
		{
			desc:       "all options - non master branch - docker",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Docker: refBool(true)},
				{Full: refBool(true)},
				{Master: refString("main")},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v1.3.0-master-d28fbce",
		},
	}

	fv := &versionfakes.FakeVersioner{}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			fv.BranchesReturns([]string{tc.branch}, nil)
			fv.CommitReturns(tc.commit, nil)
			fv.CommittishReturns(tc.committish, nil)
			fv.TagReturns(tc.tag, nil)

			v, err := version.NewVersion(fv, tc.opts...)
			if err != nil {
				t.Error("could not instanitate version:", err)
			}

			v.IncMinor()
			actual, err := v.Minor()
			if err != nil {
				t.Error("could not generate version:", err)
			}

			if actual != tc.expect {
				t.Errorf("actual(%v) generated version is not as expected(%v)", actual, tc.expect)
			}
		})
	}
}

func TestIncMinor(t *testing.T) {
	testcases := []struct {
		desc       string
		branch     string
		commit     string
		committish string
		tag        string
		opts       []*version.VersionOption
		expect     string
	}{
		{
			desc:       "default",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			expect:     "v1.2.4",
		},
		{
			desc:       "default - non master branch",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Master: refString("main")},
			},
			expect: "master",
		},
		{
			desc:       "all options",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Full: refBool(true)},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v1.2.4-master+d28fbce",
		},
		{
			desc:       "all options - non master branch",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Full: refBool(true)},
				{Master: refString("main")},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v1.2.4-master+d28fbce",
		},
		{
			desc:       "all options - non master branch - docker",
			branch:     "master",
			commit:     "d28fbcea1e82ef9bd117fa07a7664032df8437b1",
			committish: "d28fbce",
			tag:        "v1.2.3",
			opts: []*version.VersionOption{
				{Branch: refBool(true)},
				{Docker: refBool(true)},
				{Full: refBool(true)},
				{Master: refString("main")},
				{Revision: refBool(true)},
				{Semver: refBool(true)},
			},
			expect: "v1.2.4-master-d28fbce",
		},
	}

	fv := &versionfakes.FakeVersioner{}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			fv.BranchesReturns([]string{tc.branch}, nil)
			fv.CommitReturns(tc.commit, nil)
			fv.CommittishReturns(tc.committish, nil)
			fv.TagReturns(tc.tag, nil)

			v, err := version.NewVersion(fv, tc.opts...)
			if err != nil {
				t.Error("could not instanitate version:", err)
			}

			v.IncPatch()
			actual, err := v.Patch()
			if err != nil {
				t.Error("could not generate version:", err)
			}

			if actual != tc.expect {
				t.Errorf("actual(%v) generated version is not as expected(%v)", actual, tc.expect)
			}
		})
	}
}

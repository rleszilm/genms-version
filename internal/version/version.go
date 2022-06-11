package version

import (
	"bytes"
	"text/template"

	"github.com/blang/semver"
)

const (
	defaultMaster string = "master"

	defaultVersionTmplSrc = `
{{- if eq .Branch .Options.Master -}}
	v{{- .Major -}}

	{{- if or (eq .Objective "minor") (eq .Objective "patch") -}}
		.{{- .Minor -}}
	{{- end -}}

	{{- if eq .Objective "patch" -}}
		.{{- .Patch -}}
	{{- end -}}
	
	{{- if .Options.Branch -}}
		-{{- .Branch -}}
	{{- end -}}
{{- else -}}
	{{- if .Options.Semver -}}
		v{{- .Major -}}

		{{- if or (eq .Objective "minor") (eq .Objective "patch") -}}
			.{{- .Minor -}}
		{{- end -}}

		{{- if eq .Objective "patch" -}}
			.{{- .Patch -}}
		{{- end -}}
		-
	{{- end -}}
	{{- .Branch -}}
{{- end -}}

{{- if .Options.Revision -}}
	{{- if .Options.Docker -}}
		-
	{{- else -}}
		+
	{{- end -}}
	{{- .Committish -}}
{{- end -}}
`
)

var (
	defaultVersionTmpl *template.Template
)

func init() {
	tmpl, err := template.New("defaultVersionTmpl").
		Funcs(template.FuncMap{}).
		Parse(defaultVersionTmplSrc)
	if err != nil {
		panic(err)
	}

	defaultVersionTmpl = tmpl
}

// Versioner is an interface that caontains data about the version of a resource.
//counterfeiter:generate . Versioner
type Versioner interface {
	Branch() (string, error)
	Commit() (string, error)
	Committish() (string, error)
	Tag() (string, error)
}

// VersionOption defines the options that can be passed in.
type VersionOption struct {
	Branch   *bool
	Docker   *bool
	Full     *bool
	Master   *string
	Revision *bool
	Semver   *bool
}

// SetBranch sets the WithBranch flag.
func (v *VersionOption) SetBranch(b bool) {
	v.Branch = &b
}

// SetDocker sets the WithDocker flag.
func (v *VersionOption) SetDocker(b bool) {
	v.Docker = &b
}

// SetFull sets the WithFull flag.
func (v *VersionOption) SetFull(b bool) {
	v.Full = &b
}

// SetMaster sets the WithMaster flag.
func (v *VersionOption) SetMaster(s string) {
	v.Master = &s
}

// SetRevision sets the WithBuild flag.
func (v *VersionOption) SetRevision(b bool) {
	v.Revision = &b
}

// SetSemver sets the WithSemver flag.
func (v *VersionOption) SetSemver(b bool) {
	v.Semver = &b
}

// Map returns a map containing the set values.
func (v *VersionOption) Map() map[string]interface{} {
	m := map[string]interface{}{}

	if v.Branch != nil {
		m["Branch"] = *v.Branch
	}

	if v.Docker != nil {
		m["Docker"] = *v.Docker
	}

	if v.Master != nil {
		m["Master"] = *v.Master
	}

	if v.Revision != nil {
		m["Revision"] = *v.Revision
	}

	if v.Semver != nil {
		m["Semver"] = *v.Semver
	}

	return m
}

func options(vos ...*VersionOption) *VersionOption {
	opts := &VersionOption{}
	opts.SetBranch(false)
	opts.SetDocker(false)
	opts.SetMaster(defaultMaster)
	opts.SetRevision(false)
	opts.SetSemver(false)

	for _, vo := range vos {
		if vo.Branch != nil {
			opts.Branch = vo.Branch
		}
		if vo.Docker != nil {
			opts.Docker = vo.Docker
		}
		if vo.Full != nil {
			opts.Full = vo.Full
		}
		if vo.Master != nil {
			opts.Master = vo.Master
		}
		if vo.Revision != nil {
			opts.Revision = vo.Revision
		}
		if vo.Semver != nil {
			opts.Semver = vo.Semver
		}
	}
	return opts
}

// Version defines code that returns version strings.
type Version struct {
	major      uint64
	minor      uint64
	patch      uint64
	branch     string
	commit     string
	committish string
	options    *VersionOption
	tmpl       *template.Template
}

// Major returns the next major version.
func (v *Version) Major() (string, error) {
	args := v.args()
	args["Objective"] = "major"
	if v.options.Full != nil && *v.options.Full {
		args["Objective"] = "patch"
	}

	buf := &bytes.Buffer{}
	err := v.tmpl.Execute(buf, args)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Minor returns the next minor version.
func (v *Version) Minor() (string, error) {
	args := v.args()
	args["Objective"] = "minor"
	if v.options.Full != nil && *v.options.Full {
		args["Objective"] = "patch"
	}

	buf := &bytes.Buffer{}
	err := v.tmpl.Execute(buf, args)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Patch returns the next patch version.
func (v *Version) Patch() (string, error) {
	args := v.args()
	args["Objective"] = "patch"

	buf := &bytes.Buffer{}
	err := v.tmpl.Execute(buf, args)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// IncMajor increments the major version.
func (v *Version) IncMajor() {
	v.major++
	v.minor = 0
	v.patch = 0
}

// IncMinor increments the Minor version.
func (v *Version) IncMinor() {
	v.minor++
	v.patch = 0
}

// IncPatch increments the Patch version.
func (v *Version) IncPatch() {
	v.patch++
}

// WithTemplate assigns the template to use when formatting versions.
func (v *Version) WithTemplate(t *template.Template) {
	v.tmpl = t
}

func (v *Version) args() map[string]interface{} {
	return map[string]interface{}{
		"Major":      v.major,
		"Minor":      v.minor,
		"Patch":      v.patch,
		"Branch":     v.branch,
		"Commit":     v.commit,
		"Committish": v.committish,
		"Options":    v.options.Map(),
	}
}

// NewVersion returns a new Version.
func NewVersion(ver Versioner, vos ...*VersionOption) (*Version, error) {
	t, err := ver.Tag()
	if err != nil {
		return nil, err
	}

	s, err := semver.ParseTolerant(t)
	if err != nil {
		return nil, err
	}

	commit, err := ver.Commit()
	if err != nil {
		return nil, err
	}

	branch, err := ver.Branch()
	if err != nil {
		return nil, err
	}

	committish, err := ver.Committish()
	if err != nil {
		return nil, err
	}

	v := &Version{
		major:      s.Major,
		minor:      s.Minor,
		patch:      s.Patch,
		branch:     branch,
		commit:     commit,
		committish: committish,
		options:    options(vos...),
		tmpl:       defaultVersionTmpl,
	}

	return v, nil
}

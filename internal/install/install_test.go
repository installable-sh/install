package install

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/pem"
	"strings"
	"testing"
	"time"

	"github.com/installable-sh/lib/certs"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want *Install
	}{
		{
			name: "empty args",
			args: []string{},
			want: &Install{},
		},
		{
			name: "help flag",
			args: []string{"--help"},
			want: &Install{ShowHelp: true},
		},
		{
			name: "help short flag",
			args: []string{"-h"},
			want: &Install{ShowHelp: true},
		},
		{
			name: "version flag",
			args: []string{"--version"},
			want: &Install{ShowVersion: true},
		},
		{
			name: "version short flag",
			args: []string{"-v"},
			want: &Install{ShowVersion: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args)

			if got.ShowHelp != tt.want.ShowHelp {
				t.Errorf("ShowHelp = %v, want %v", got.ShowHelp, tt.want.ShowHelp)
			}
			if got.ShowVersion != tt.want.ShowVersion {
				t.Errorf("ShowVersion = %v, want %v", got.ShowVersion, tt.want.ShowVersion)
			}
		})
	}
}

func TestExec_Help(t *testing.T) {
	var stdout bytes.Buffer
	cmd := &Install{
		ShowHelp: true,
		Stdout:   &stdout,
		Stderr:   &bytes.Buffer{},
	}

	err := cmd.Exec(context.Background())
	if err != nil {
		t.Errorf("Exec() error = %v", err)
	}

	if !strings.Contains(stdout.String(), "usage:") {
		t.Errorf("Exec() stdout = %q, want usage info", stdout.String())
	}
}

func TestExec_Version(t *testing.T) {
	var stdout bytes.Buffer
	cmd := &Install{
		ShowVersion: true,
		Stdout:      &stdout,
		Stderr:      &bytes.Buffer{},
	}

	err := cmd.Exec(context.Background())
	if err != nil {
		t.Errorf("Exec() error = %v", err)
	}
}

func TestExec_NotImplemented(t *testing.T) {
	var stderr bytes.Buffer
	cmd := &Install{
		Stdout: &bytes.Buffer{},
		Stderr: &stderr,
	}

	err := cmd.Exec(context.Background())
	if err == nil {
		t.Error("Exec() expected error for not implemented")
	}

	if !strings.Contains(stderr.String(), "coming soon") {
		t.Errorf("Exec() stderr = %q, want 'coming soon'", stderr.String())
	}
}

func TestCerts_NonEmptyAndFresh(t *testing.T) {
	if len(certs.CACerts) == 0 {
		t.Fatal("CACerts is empty")
	}

	// Parse all PEM blocks and check for valid, non-expired certs
	var validCerts int
	var freshCerts int
	now := time.Now()
	data := certs.CACerts

	for {
		block, rest := pem.Decode(data)
		if block == nil {
			break
		}
		data = rest

		if block.Type != "CERTIFICATE" {
			continue
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			continue
		}
		validCerts++

		if now.Before(cert.NotAfter) && now.After(cert.NotBefore) {
			freshCerts++
		}
	}

	if validCerts == 0 {
		t.Fatal("no valid certificates found in CACerts")
	}

	if freshCerts == 0 {
		t.Fatal("no fresh (non-expired) certificates found in CACerts")
	}

	t.Logf("found %d valid certificates, %d are fresh", validCerts, freshCerts)
}

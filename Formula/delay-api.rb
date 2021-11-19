# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class DelayApi < Formula
  desc "simple wait& response api for timeout test"
  homepage "https://github.com/applegreengrape/delay-api"
  version "0.0.2"
  depends_on :macos

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/applegreengrape/delay-api/releases/download/v0.0.2/delay-api_0.0.2_darwin_amd64.zip"
      sha256 "afd4c63c09a64f633de447c2eccabe99c1af9af140e26e0a3ed77851079b22dc"

      def install
        bin.install "wait"
      end
    end
  end
end

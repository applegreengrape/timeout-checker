# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class DelayApi < Formula
  desc "simple wait& response api for timeout test"
  homepage "https://github.com/applegreengrape/delay-api"
  version "0.0.1"
  depends_on :macos

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/applegreengrape/delay-api/releases/download/v0.0.1/delay-api_0.0.1_darwin_amd64.zip"
      sha256 "788bf7c5fe4c1a32af08918e5b2882db50883c264e5a7acab13848934a27b139"

      def install
        bin.install "wait"
      end
    end
  end
end